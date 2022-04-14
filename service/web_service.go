package service

import (
	"bytes"
	"log"
	"net/url"
	"strings"
	"webscraper/config"
	"webscraper/model"
	"webscraper/util"

	"golang.org/x/net/html"
)

type WebPageService struct {
	config *config.Config
}

func NewWebPageService(c *config.Config) *WebPageService {
	return &WebPageService{config: c}
}

func (h *WebPageService) Analyze(uri string) (*model.HtmlDocument, error) {
	doc := model.HtmlDocument{}

	version, err := util.DetectFromURL(uri)
	if err != nil {
		log.Printf("HTML version detection error occured : %s", err.Error())
		return nil, err
	}

	doc.HTMLVersion = version

	if url, err := url.Parse(uri); err == nil {
		doc.Host = url.Host
	} else {
		return nil, err
	}

	docBytes, err := util.ReadBytes(uri)
	if err != nil {
		return nil, err
	}

	parsedDoc, _ := html.Parse(bytes.NewBuffer(docBytes))
	return crawl(parsedDoc, &doc)

}

func crawl(node *html.Node, doc *model.HtmlDocument) (*model.HtmlDocument, error) {
	switch {
	case node.Parent != nil && node.Parent.Type == html.ElementNode:
		switch strings.ToLower(node.Parent.Data) {
		case "title":
			doc.Title = node.Data
		case "h1":
			doc.H1Count++
		case "h2":
			doc.H2Count++
		case "h3":
			doc.H3Count++
		case "h4":
			doc.H4Count++
		case "h5":
			doc.H5Count++
		case "h6":
			doc.H6Count++
		case "a":
			for _, attr := range node.Parent.Attr {
				if strings.ToLower(attr.Key) == "href" && attr.Val != "" {
					if u, err := url.Parse(attr.Val); err == nil {
						tempLink := &model.Link{Url: attr.Val, Accessible: util.IsReachable(attr.Val)}
						if u.Host != "" && u.Host != doc.Host {
							// util.IsReachable(attr.Val)
							doc.ExternalLinks = append(doc.ExternalLinks, *tempLink)
						} else {
							tempLink.Accessible = true
							doc.InternalLinks = append(doc.InternalLinks, *tempLink)
						}
						break
					}
				}
			}
		}
	case node.Type == html.ElementNode && node.Data == "input":
		for _, attr := range node.Attr {
			if strings.ToLower(attr.Val) == "password" {
				doc.LoginExist = true
				break
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		crawl(child, doc)
	}
	return doc, nil
}
