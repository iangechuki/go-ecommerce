
import { ProductCard } from "./product-card";
import { Product } from "@/types";
interface ProductGridProps {
    products:Product[]
}
export function ProductGrid({products}:ProductGridProps){
    return (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
            {products.map((product) => (
                <ProductCard key={product.id} product={product}/>
            ))}
        </div>
    )
}