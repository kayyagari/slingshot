package jnlp

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func ParseFromFile(jnlpPath string) (*Jnlp, error) {
	f, err := os.Open(jnlpPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return _parse(f)
}

func _parse(r io.Reader) (*Jnlp, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var j Jnlp
	err = xml.Unmarshal(data, &j)
	if err != nil {
		return nil, err
	}

	return &j, nil
}

func ParseFromUrl(jnlpUrl string) (*Jnlp, error) {
	u, err := url.Parse(jnlpUrl)
	if err != nil {
		return nil, err
	}
	cl := http.Client{}
	if u.Scheme == "https" {
		tlsConf := &tls.Config{InsecureSkipVerify: true}
		cl.Transport = &http.Transport{TLSClientConfig: tlsConf}
	}

	return _parseFromUrl(cl, jnlpUrl)
}

func _parseFromUrl(cl http.Client, jnlpUrl string) (*Jnlp, error) {
	req, _ := http.NewRequest(http.MethodGet, jnlpUrl, nil)
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	tlp, err := _parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return tlp, nil
}

func Download(jnlp *Jnlp) (string, error) {
	u, _ := url.Parse(jnlp.Codebase)
	p := filepath.Join("/tmp", jnlp.Version, u.Host)
	fi, _ := os.Stat(p)
	if fi != nil {
		os.RemoveAll(p)
	}

	err := os.MkdirAll(p, 0744)
	if err != nil {
		return "", err
	}

	cl := http.Client{}
	if u.Scheme == "https" {
		tlsConf := &tls.Config{InsecureSkipVerify: true}
		cl.Transport = &http.Transport{TLSClientConfig: tlsConf}
	}

	err = downloadJars(p, jnlp, cl)

	return p, err
}

func downloadJars(baseDir string, jnlp *Jnlp, cl http.Client) error {
	for _, r := range jnlp.Resources {
		for _, j := range r.Jar {
			p := jnlp.Codebase + "/" + j.Href
			resp, err := cl.Get(p)
			if err != nil {
				return err
			}
			fmt.Printf("downloading jar %s\n", p)
			pos := strings.LastIndex(p, "/")
			name := p[pos+1:]
			name = filepath.Join(baseDir, name)
			data, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			err = ioutil.WriteFile(name, data, 0644)
			if err != nil {
				return err
			}
		}

		for _, e := range r.Extension {
			p := jnlp.Codebase + "/" + e.Href
			eJnlp, err := _parseFromUrl(cl, p)
			if err != nil {
				return err
			}
			pos := strings.LastIndex(p, "/")
			eJnlp.Codebase = p[:pos]
			err = downloadJars(baseDir, eJnlp, cl)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func ParseAndDownload(jnlpUrl string) (string, *Jnlp, error) {
	j, err := ParseFromUrl(jnlpUrl)
	if err != nil {
		return "", nil, err
	}

	dirPath, err := Download(j)
	if err != nil {
		return "", nil, err
	}

	return dirPath, j, nil
}
