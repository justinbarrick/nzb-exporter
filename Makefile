deployment.yaml: main.go config/template.yaml
	ko resolve -f config/template.yaml > deployment.yaml
