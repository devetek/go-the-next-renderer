const templateGenerator = (html: string): string => {
    // lexer / parser
    return html.replaceAll("Ikan Cuek", "{{ .Price }}");
    return html
}

export default templateGenerator;