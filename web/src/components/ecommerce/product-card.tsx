"use client"
import Image from "next/image";
import { Badge } from "../ui/badge";
import { Star } from "lucide-react";
import { Button } from "../ui/button";
import Link from "next/link";
import { Product } from "@/types";
import { useState } from "react";
import { useCart } from "@/lib/stores/cart-store";
import {toast} from "sonner"
interface ProductCardProps {
    product:Product
}
export function ProductCard({ product }: ProductCardProps) {
  const [quantity,setQuantity] = useState(1);
  const {addItem} = useCart()
  const handleAddToCart = ()=>{
   console.log(quantity)
   addItem(product,quantity)
   toast.success("Added to cart",{
    description:`${quantity} x ${product.name}`,
    action:{
      label:"View Cart",
      onClick:()=> (window.location.href="/cart")
    }
   })
   setQuantity(1)
  }
  return (
    <div className="overflow-hidden border rounded-xl hover:shadow-lg transition-shadow">
      <div className="relative aspect-square bg-muted">
        <Image
          src={product.images[0]}
          alt={product.name}
          fill
          className="object-cover"
        />
        {/* <Badge variant={"destructive"} className="absolute top-2 left-2">
          -20%
        </Badge> */}
         <Link 
          href={`/products/${product.id}`} 
          className="absolute inset-0 z-10"
          aria-label={`View ${product.name} details`}
        >
          <span className="sr-only">View Product</span>
        </Link>
      </div>
      <div className="p-4 space-y-4">
        <div className="space-y-2">
          <h3 className="font-semibold">{product.name}</h3>
          <div className="flex items-center gap-2">
            <span className="text-lg font-bold">${product.price}</span>
            {product.originalPrice && (
              <span className="line-through text-muted-foreground">
                ${product.originalPrice}
              </span>
            )}
          </div>
        </div>

        <div className="flex gap-2">
          <Button
            variant="outline"
            className="flex-1"
            asChild
          >
            <Link href={`/products/${product.id}`}>
              View Product
            </Link>
          </Button>
          <Button
            className="flex-1"
            onClick={handleAddToCart}
          >
            Add to Cart
          </Button>
        </div>
        </div>
    </div>
  );
}
