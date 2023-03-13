import React, { FC } from "react";
import RelatedItem from "./Item";


interface RelatedProps  {
    relateds: Product[];
}

const RelatedComponent: FC<RelatedProps> = ({
  relateds,
}: RelatedProps): JSX.Element => {
  return (
      <div data-tmpl="go:loop">
        <ul>
          {!!relateds.length &&
            relateds.map((product: Product) => {
              return (<RelatedItem key={product.id} related={product}/>);
            })}
        </ul>
      </div>
  );
};

export default RelatedComponent;
