import React, { FC } from "react";
import HTMLReactParser from "html-react-parser";


interface RelatedItemProps  {
    related: Product;
}

const RelatedItemComponent: FC<RelatedItemProps> = ({
  related,
}: RelatedItemProps): JSX.Element => {
  return (
    <li key={related.id}>
        <div>
            <h4 data-tmpl="go:related.Name">{related.name}</h4>
        </div>
        <div data-tmpl="go:related.Price">{related.price}</div>
        {related.discount && <div data-tmpl="go:related.Discount">{related.discount}</div>}
    </li>
  );
};

export default RelatedItemComponent;
