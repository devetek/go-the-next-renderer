import React, { FC, Fragment } from "react";
import HTMLReactParser from "html-react-parser";
import { GetServerSideProps, GetServerSidePropsContext } from "next";
import fetchData from "../../utils/fetch";
import RelatedComponent from "@/components/Related";

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
  product,
  related,
}: ProductProps): JSX.Element => {
  return (
    <main>
      <div>
        <h1 data-tmpl="go:.Name">{product.name}</h1>
      </div>
      <div>
        <h3 data-tmpl="go:.Price">{product.price}</h3>
      </div>
      {product.discount && (
        <div>
          <h3 data-tmpl="go:.Discount">{product.discount}</h3>
        </div>
      )}
      <div data-tmpl="go:.Description">{product.description}</div>
      {!!related.length && <RelatedComponent relateds={related}/>}
    </main>
  );
};

export default ProductPage;
