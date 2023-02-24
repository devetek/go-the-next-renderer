import React, { FC, Fragment } from "react";
import HTMLReactParser from "html-react-parser";
import { GetServerSideProps, GetServerSidePropsContext } from "next";
import fetchData from "../utils/fetch";

export const getServerSideProps: GetServerSideProps = async ({
  req,
}: GetServerSidePropsContext) => {
  const ua = req?.headers?.["user-agent"] || "";
  const isTemplateConsumer = ua?.includes("wpe-bot-aggregator");
  const pickTemplate = isTemplateConsumer ? "topics" : "";
  const data = await fetchData(
    pickTemplate,
    "https://mocki.io/v1/15b1a310-fbe6-453a-9fcb-57ec3e240cb1"
  );

  return {
    props: {
      tConsumer: isTemplateConsumer,
      ...data,
    }, // will be passed to the page component as props
  };
};

type LoopProps = {
  tConsumer: boolean;
  topics: string[];
};

const LoopPage: FC<any> = ({ tConsumer, topics }: LoopProps): JSX.Element => {
  return (
    <main>
      {!!topics.length &&
        topics.map((topic: string) => {
          return tConsumer ? (
            HTMLReactParser(topic)
          ) : (
            <p key={topic}>{topic}</p>
          );
        })}
    </main>
  );
};

export default LoopPage;
