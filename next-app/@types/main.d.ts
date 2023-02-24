interface Category {
    id: number;
    name: string;
}

interface Shop {
    id: number;
    name: string;
    description: string;
}

interface Product {
    id: number;
    name: string;
    category: Category;
    price: string;
    discount?: string;
    price_raw: number;
    description: string;
}

interface ResponseShopDetail {
    shop: Shop;
    products: Product[];
}

interface ResponseProductDetail {
    product: Product;
    related: Product[];
}