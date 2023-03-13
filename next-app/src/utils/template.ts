const templateGenerator = (html: string): string => {
    // lexer / parser
    return html.replaceAll("Rp. 10.000,-", "{{ .Price }}");
}

export default templateGenerator;