// components/ecommerce/add-to-cart-button.tsx
"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";
import { Plus, Minus } from "lucide-react";
import { useCart } from "@/lib/stores/cart-store";

export function AddToCartButton({ 
  product,
  variant = "default"
}: {
  product: any;
  variant?: "default" | "large";
}) {
  const [quantity, setQuantity] = useState(1);
  const { addItem } = useCart();

  const handleAddToCart = () => {
    addItem(product, quantity);
    setQuantity(1);
  };

  return (
    <div className={`flex ${variant === "large" ? "gap-4" : "gap-2"}`}>
      <div className="flex items-center border rounded-lg">
        <Button
          variant="ghost"
          size="icon"
          onClick={() => setQuantity(Math.max(1, quantity - 1))}
          className="h-8 w-8 rounded-r-none"
        >
          <Minus className="h-4 w-4" />
        </Button>
        
        <Input
          type="number"
          value={quantity}
          min="1"
          max={product.stock}
          onChange={(e) => setQuantity(Math.min(product.stock, Math.max(1, Number(e.target.value))))}
          className="w-12 h-8 text-center border-0 rounded-none [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
        />
        
        <Button
          variant="ghost"
          size="icon"
          onClick={() => setQuantity(Math.min(product.stock, quantity + 1))}
          className="h-8 w-8 rounded-l-none"
        >
          <Plus className="h-4 w-4" />
        </Button>
      </div>
      
      <Button 
        onClick={handleAddToCart}
        className={variant === "large" ? "px-8 py-4 text-lg" : ""}
      >
        Add to Cart
      </Button>
    </div>
  );
}