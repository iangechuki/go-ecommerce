import { ProductGrid } from "@/components/ecommerce/product-grid";
import { Button } from "@/components/ui/button";
import { products } from "@/constants/products";
import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <div className="space-y-12">
      {/* Hero */}
      <section className="relative bg-gradient-to-r from-primary/10 to-secondary/10">
        <div className="container flex flex-col md:flex-row items-center gap-8 py-16 px-2">
          <div className="md:w-1/2 space-y-6">
            <h1 className="text-4xl md:text-5xl font-bold tracking-tight">Discover Your Perfect Tech Gear</h1>
            <p className="text-lg text-muted-foreground">
              Explore cutting edge electronic goods and accessories at
              unbeatable prices.Free shipping on orders over 100
            </p>
            <Button size="lg" asChild>
              <Link href="/products">Shop Now</Link>
            </Button>
          </div>
          <div className="aspect-video md:w-1/2 w-full relative">
            <Image
              src="https://plus.unsplash.com/premium_photo-1676998931123-75789162f170?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8aGVybyUyMGltYWdlfGVufDB8fDB8fHww"
              fill
              alt="Electric collection"
              className="rounded-xl object-cover"
            />
          </div>
        </div>
      </section>
      {/* Featured Categories*/}
      <section className="container p-2">
        <h2 className="text-3xl font-bold mb-8">Shop by category</h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <Link href="/cat" className="group relative block overflow-hidden rounded-xl">
          <div className="relative aspect-square bg-muted">
                <Image
                  src={`/images/android-img.png`}
                  alt={"cat-name"}
                  fill
                  className="object-cover transition-transform group-hover:scale-105"
                />
              </div>
              <div className="absolute inset-0 bg-black/30 flex items-center justify-center">
                <h3 className="text-2xl font-bold text-white">
                  {"cat-name"}
                </h3>
              </div>
          </Link>
          <Link href="/cat" className="group relative block overflow-hidden rounded-xl">
          <div className="relative aspect-square bg-muted">
                <Image
                  src={`/images/earbuds-1.jpg`}
                  alt={"cat-name"}
                  fill
                  className="object-cover transition-transform group-hover:scale-105"
                />
              </div>
              <div className="absolute inset-0 bg-black/30 flex items-center justify-center">
                <h3 className="text-2xl font-bold text-white">
                  {"cat-name"}
                </h3>
              </div>
          </Link>
          <Link href="/cat" className="group relative block overflow-hidden rounded-xl">
          <div className="relative aspect-square bg-muted">
                <Image
                  src={`/images/android-img.png`}
                  alt={"cat-name"}
                  fill
                  className="object-cover transition-transform group-hover:scale-105"
                />
              </div>
              <div className="absolute inset-0 bg-black/30 flex items-center justify-center">
                <h3 className="text-2xl font-bold text-white">
                  {"cat-name"}
                </h3>
              </div>
          </Link>
        </div>
      </section>
      {/* Featured Products */}
      <section className="container p-2">
        <div className="flex justify-between items-center mb-8">
          <h2 className="text-3xl font-bold">Popular Products</h2>
          <Button variant={"link"} asChild>
            <Link href="/products">View All Products</Link>
          </Button>
        </div>
        <ProductGrid products={products}/>
      </section>
      {/* Promo Banner */}
      <section className="bg-primary text-primary-foreground py-16">
        <div className="container text-center space-y-4">
          <h2 className="text-3xl font-bold">Summer Sale!</h2>
          <p className="text-lg">Up to 50% off selected items</p>
          <Button variant="secondary" asChild>
            <Link href="/products">Shop Sale</Link>
          </Button>
        </div>
      </section>
    </div>
  );
}
