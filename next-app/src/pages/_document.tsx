import Document, { DocumentContext, DocumentInitialProps, Head, Html, Main, NextScript } from 'next/document'
// import templateGenerator from '@/utils/template'

class RendererDocument extends Document {
  static async getInitialProps(
    ctx: DocumentContext
  ): Promise<DocumentInitialProps> {
    const initialProps = await Document.getInitialProps(ctx)


    // Here, we create our template for target engine (Golang / Rust / V-Lang).
    // Modify HTML with contract parser / lexter (we cannot modify initialProps.html) just like this!
    // initialProps.html =  templateGenerator(initialProps.html)

    return initialProps
  }

  render() {
    return (
      <Html>
        <Head />
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    )
  }
}

export default RendererDocument