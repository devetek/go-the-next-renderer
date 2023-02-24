import React, { FC, Fragment } from "react";
import HTMLReactParser from "html-react-parser";
import { GetServerSideProps, GetServerSidePropsContext } from "next";
import fetchData from "../../utils/fetch";

const getLinkAPI = (target: string): string => {
  const apiMap: Record<string, string> = {
    a: "https://mocki.io/v1/75805603-5327-4a4e-af47-af8307988c35",
    b: "https://mocki.io/v1/dc40078b-3d11-426c-9f61-c41c2e2428de",
    default: "https://mocki.io/v1/dcb503d6-671a-4ecd-b5ee-443dbd8806db",
  };
  const selectedData = apiMap[target];

  if (!selectedData) {
    return apiMap["default"];
  }

  return selectedData;
};

export const getServerSideProps: GetServerSideProps = async ({
  req,
  params,
}: GetServerSidePropsContext) => {
  const name = params?.name || "";
  const ua = req?.headers?.["user-agent"] || "";
  const isTemplateConsumer = ua?.includes("wpe-bot-aggregator");
  const pickTemplate = isTemplateConsumer ? "product-detail" : "";
  const apiLink = getLinkAPI(name.toString());
  const data: ResponseProductDetail = await fetchData(pickTemplate, apiLink);

  return {
    props: {
      tConsumer: isTemplateConsumer,
      ...data,
    }, // will be passed to the page component as props
  };
};

interface ProductProps extends ResponseProductDetail {
  tConsumer: boolean;
}

const ProductPage: FC<ProductProps> = ({
  tConsumer,
  product,
  related,
}: ProductProps): JSX.Element => {
  return (
    <main>
      <div>
        <h1>{product.name}</h1>
      </div>
      <div>
        <h3>{product.price}</h3>
      </div>
      {tConsumer
        ? HTMLReactParser(product?.discount || "")
        : product.discount && (
            <div>
              <h3>{product.discount}</h3>
            </div>
          )}
      <div>{HTMLReactParser(product.description)}</div>
      <div>
        <ul>
          {tConsumer &&
            related.map((product: Product) => HTMLReactParser(product.name))}
          {!tConsumer &&
            !!related.length &&
            related.map((product: Product) => {
              return (
                <li key={product.id}>
                  <div>
                    <h4>{product.name}</h4>
                  </div>
                  <div>{product.price}</div>
                  {tConsumer
                    ? HTMLReactParser(product?.discount || "")
                    : product.discount && <div>{product.discount}</div>}
                </li>
              );
            })}
        </ul>
      </div>
    </main>
  );
};

export default ProductPage;
