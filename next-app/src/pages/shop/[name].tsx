import React, { FC, Fragment } from "react";
import HTMLReactParser from "html-react-parser";
import { GetServerSideProps, GetServerSidePropsContext } from "next";
import fetchData from "../../utils/fetch";

const getLinkAPI = (target: string): string => {
  const apiMap: Record<string, string> = {
    a: "https://mocki.io/v1/61674e3c-675f-4cfe-a38d-cf676bffbda1",
    b: "https://mocki.io/v1/07ac5dc4-ac87-4bf7-8053-f61de366d136",
    default: "https://mocki.io/v1/d1f33761-e36c-410f-8cfb-9014caf6f3a1",
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
  const pickTemplate = isTemplateConsumer ? "shop" : "";
  const apiLink = getLinkAPI(name.toString());
  const data: ResponseShopDetail = await fetchData(pickTemplate, apiLink);

  return {
    props: {
      tConsumer: isTemplateConsumer,
      ...data,
    }, // will be passed to the page component as props
  };
};

interface ShopProps extends ResponseShopDetail {
  tConsumer: boolean;
}

const ShopPage: FC<any> = ({
  tConsumer,
  shop,
  products,
}: ShopProps): JSX.Element => {
  return (
    <main>
      <div>
        <h1>{shop.name}</h1>
      </div>
      <div>{HTMLReactParser(shop.description)}</div>
      <div>
        <ul>
          {products &&
            !!products.length &&
            products.map((product: Product) => {
              return tConsumer ? (
                HTMLReactParser(product.name)
              ) : (
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

export default ShopPage;
