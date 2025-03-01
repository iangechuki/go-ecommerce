"use client";
import { ProductImages } from "@/components/ecommerce/product-images";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { products } from "@/constants/products";
import { useCart } from "@/lib/stores/cart-store";
import { Minus, Plus } from "lucide-react";
import Link from "next/link";
import { useState } from "react";
import { toast } from "sonner";

export default function ProductPage({ id }: { id: string }) {
  const [quantity, setQuantity] = useState(1);
  const { addItem } = useCart();
  const product = products.find((product) => product.id === id);
  const handleAddToCart = () => {
    if (!product) return;

    addItem(product, quantity);
    toast.success("Added to cart", {
      description: `${quantity} x ${product.name}`,
      action: {
        label: "View Cart",
        onClick: () => (window.location.href = "/cart"),
      },
    });
  };
  if (!product)
    return (
      <div className="container py-8 text-center">
        <h1 className="text-2xl font-bold mb-4">Product not found</h1>
        <Button asChild>
          <Link href="/products">Browse Products</Link>
        </Button>
      </div>
    );
  return (
    <div className="container px-4 py-8">
      <div className="grid md:grid-cols-2 gap-8">
        <ProductImages images={product.images} />
        <div className="space-y-6">
          <div>
            <h1 className="text-3xl font-bold">{product.name}</h1>
            <div className="flex items-center gap-4">
              <span className="text-2xl font-bold">${product.price}</span>
              {product.originalPrice && (
                <span className="line-through text-muted-foreground">
                  ${product.originalPrice}
                </span>
              )}
            </div>
            <p className="text-sm text-muted-foreground">
              {product.stock > 0
                ? `${product.stock} items in stock`
                : "Out of stock"}
            </p>
          </div>
          <div className="flex items-center gap-4">
            <div className="flex items-center border rounded-md">
              <Button
                variant="ghost"
                size="sm"
                disabled={quantity <= 1}
                className="h-9 px-2"
                onClick={() => setQuantity(Math.max(1, quantity - 1))}
              >
                <Minus className="h-4 w-4" />
              </Button>
              <Input
                type="number"
                className="w-16 h-9 text-center border-0 [appearance:textfield]"
                value={quantity}
                min={1}
                max={product.stock}
                onChange={(e) => {
                  const value = Math.max(
                    1,
                    Math.min(product.stock, Number(e.target.value))
                  );
                  setQuantity(value);
                }}
              />
              <Button
                variant="ghost"
                size="sm"
                className="h-9 px-2"
                onClick={() =>
                  setQuantity(Math.min(product.stock, quantity + 1))
                }
                disabled={quantity >= product.stock}
              >
                <Plus className="h-4 w-4" />
              </Button>
            </div>
            <Button
              className="flex-1 h-9"
              onClick={handleAddToCart}
              disabled={product.stock === 0}
            >
              Add to Cart
            </Button>
          </div>
          <Tabs defaultValue="description">
            <TabsList>
              <TabsTrigger value="description">Description</TabsTrigger>
              <TabsTrigger value="specifications">Specifications</TabsTrigger>
              <TabsTrigger value="reviews">
                Reviews ({product.reviews})
              </TabsTrigger>
            </TabsList>

            <TabsContent value="description" className="pt-4">
              <p className="text-muted-foreground">{product.description}</p>
            </TabsContent>

            <TabsContent value="specifications" className="pt-4">
              <div className="space-y-2">
                <div className="flex justify-between border-b py-2">
                  <span>Category</span>
                  <span className="text-muted-foreground capitalize">
                    {product.category}
                  </span>
                </div>
                <div className="flex justify-between border-b py-2">
                  <span>Rating</span>
                  <span className="text-muted-foreground">
                    {product.rating}/5
                  </span>
                </div>
              </div>
            </TabsContent>

            <TabsContent value="reviews" className="pt-4">
              {/* Add review component here */}
              <p className="text-muted-foreground">
                Customer reviews will be displayed here
              </p>
            </TabsContent>
          </Tabs>
        </div>
      </div>
    </div>
  );
}
