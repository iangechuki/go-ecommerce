"use client";

import { ProductFilter } from "@/components/ecommerce/product-filter";
import { ProductGrid } from "@/components/ecommerce/product-grid";
import { Button } from "@/components/ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import { products } from "@/constants/products";

import { ChevronDown, SlidersHorizontal } from "lucide-react";
import { useState } from "react";

export default function ProductsPage() {
  const [searchQuery, setSearchQuery] = useState("");
  const [selectedCategories, setSelectedCategories] = useState<string[]>([]);
  const [priceRange, setPriceRange] = useState<[number, number]>([0, 1000]);
  const [sortBy, setSortBy] = useState("");
  const filteredProducts = products
    .filter((product) =>
      product.name.toLowerCase().includes(searchQuery.toLowerCase()) || 
      product.description.toLocaleLowerCase().includes(searchQuery.toLocaleLowerCase())
    )
    .filter(
      (product) =>
        product.price >= priceRange[0] && product.price <= priceRange[1]
    )
    .filter(
      (product) =>
        selectedCategories.length === 0 ||
        selectedCategories.includes(product.category)
    ).sort((a, b) => {
      switch (sortBy) {
        
        case "rating":
          return b.rating - a.rating;
        case "price-low-high":
          return a.price - b.price;
        case "price-high-low":
          return b.price - a.price;
        default:
          return 0;
      }
    });
  return (
    <div className="px-4 py-8">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <div className="w-full md:w-1/3">
          <Input
            placeholder="Search products..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
          />
        </div>
        <div className="flex items-center gap-4">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
            <Button variant={"outline"}>
              Sort By: {sortByLabels[sortBy]}
              <ChevronDown className="ml-2 h-4 w-4" />
            </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem onClick={() => setSortBy("featured")}>
                Featured
              </DropdownMenuItem>
              <DropdownMenuItem onClick={() => setSortBy("price-low-high")}>
                Price: Low to High
              </DropdownMenuItem>
              <DropdownMenuItem onClick={() => setSortBy("price-high-low")}>
                Price: High to Low
              </DropdownMenuItem>
              <DropdownMenuItem onClick={() => setSortBy("rating")}>
                Rating
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
          <Button>
            <SlidersHorizontal className="mr-2 h-4 w-4" />
          </Button>
        </div>
      </div>
      <div className="grid grid-cols-1 lg:grid-cols-4 gap-8">
        <div className="lg:col-span-1">
          <ProductFilter
            priceRange={priceRange}
            onPriceChange={setPriceRange}
            selectedCategories={selectedCategories}
            onCategoryChange={setSelectedCategories}
          />
        </div>
        <div className="lg:col-span-3">
          {filteredProducts.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-muted-foreground">No Products found</p>
              <Button
                variant="ghost"
                className="mt-4"
                onClick={() => {
                  setSelectedCategories([]);
                  setPriceRange([0, 1000]);
                  setSearchQuery("");
                }}
              >
                Clear Filters
              </Button>
            </div>
          ) : (
            <>
              <p className="text-muted-foreground mb-4">
                Showing {filteredProducts.length} results
              </p>
              <ProductGrid products={filteredProducts} />
            </>
          )}
        </div>
      </div>
    </div>
  );
}

const sortByLabels: Record<string, string> = {
  featured: "Featured",
  "price-low-high": "Price:Low to High",
  "price-high-low": "Price:High to Low",
  rating: "Rating",
};
