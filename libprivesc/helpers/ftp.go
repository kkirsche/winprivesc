package helpers

import (
  "fmt"
  "text/template"
  "os"
)

const ftpTmpl = `open {{ .IP }} {{ .Port }}
{{ .Username }}
{{ .Password }}
{{- if .Binary -}}
binary
{{- end }}
{{- range .BinaryFiles }}
GET {{ .Name }}
{{- end }}
{{- if .Ascii }}
ascii
{{- end }}
{{- range .AsciiFiles }}
GET {{ .Name }}
{{- end }}
bye
`

type FTP struct {
  Path string

  IP string
  Port int
  Username string
  Password string

  Binary bool
  BinaryStrFiles []string
  BinaryFiles []File

  Ascii bool
  AsciiFilesStr []string
  AsciiFiles []File
}

type File struct {
  Name string
}


func (fh *FTP) Setup() bool {
  if fh.IP == "" {
    fmt.Println("[!] No IP address. Exiting...")
    return false
  }

  if fh.Port <= 0 {
    fmt.Println("[!] Invalid port. Exiting...")
    return false
  }

  if fh.Username == "" {
    fmt.Println("[!] No username. Exiting...")
    return false
  }

  if len(fh.AsciiFilesStr) <= 0 && len(fh.BinaryStrFiles) <= 0 {
    fmt.Println("[!] No files to transfer. Exiting...")
    return false
  }

  if len(fh.AsciiFilesStr) > 0 {
    fh.Ascii = true
    for _, f := range fh.AsciiFilesStr {
      fh.AsciiFiles = append(fh.AsciiFiles, File{Name:f})
    }
  }

  if len(fh.BinaryStrFiles) > 0 {
    fh.Binary = true
    for _, f := range fh.BinaryStrFiles {
      fh.BinaryFiles = append(fh.BinaryFiles, File{Name:f})
    }
  }

  return true
}

func (fh *FTP) OutputFile() {
  t := template.Must(template.New("ftp").Parse(ftpTmpl))
  if fh.Path == "" {
    err := t.Execute(os.Stdout, fh)
    if err != nil {
      fmt.Println("[!] Error while executing template")
      fmt.Println(fmt.Sprintf("\t%s", err))
      return
    }
  } else {
    f, err := os.OpenFile(fh.Path, os.O_RDWR|os.O_CREATE, 0764)
    if err != nil {
      fmt.Println("[!] Failed to open file")
      fmt.Println(fmt.Sprintf("\t%s", err))
      return
    }
    err = t.Execute(f, fh)
    if err != nil {
      fmt.Println("[!] Error while executing template")
      fmt.Println(fmt.Sprintf("\t%s", err))
      return
    }
  }
}
