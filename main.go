package main

import (
	"bufio"
	"fmt"
	"os"
)

type XmlShortcut struct {
	Name, Value, Description, filePath string
}

func newXmlShortcut(name, value, description, filePath string) *XmlShortcut {
	return &XmlShortcut{
		Name:        name,
		Value:       value,
		Description: description,
		filePath:    filePath,
	}
}

func (s *XmlShortcut) creatingXML() string {
	return fmt.Sprintf(
		"<template name=\"%s\" value=\"%s\" description=\"%s\" toReformat=\"true\" toShortenFQNames=\"true\">\n"+
			"    <context>\n"+
			"        <option name=\"JAVA_CODE\" value=\"true\" />\n"+
			"    </context>\n"+
			"</template>\n", s.Name, s.Value, s.Description)
}

func (s *XmlShortcut) openXMLFileAndWrite() error {
	file, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	xmlContent := s.creatingXML()

	// Adiciona o conteúdo ao arquivo
	_, err = writer.WriteString(xmlContent)
	if err != nil {
		return err
	}

	// Certifique-se de que todos os dados foram gravados no arquivo
	return writer.Flush()
}

func main() {
	filePath := "./src/main/resources/liveTemplates/ShortCuts.xml"

	db := newXmlShortcut("fl", "Float", "Shortcut for Float wrapper class", filePath)

	err := db.openXMLFileAndWrite()
	if err != nil {
		fmt.Printf("Erro ao escrever no arquivo: %v\n", err)
	} else {
		fmt.Println("Conteúdo adicionado com sucesso ao arquivo.")
	}
}
