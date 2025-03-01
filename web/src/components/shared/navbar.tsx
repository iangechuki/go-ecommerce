import Link from "next/link";
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
} from "../ui/navigation-menu";
import { cn } from "@/lib/utils";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { CartSheet } from "../ecommerce/cart-sheet";
import { UserMenu } from "./user-menu";

export function Navbar() {
  const isLoading = false;
  const isAuthenticated = false;
  return (
    <header className="sticky top-0 z-50 w-full border-b bg-background">
      <div className="conteiner flex h-16 items-center gap-4">
        <Link href="/" className="font-bold text-xl">
          ShopSphere
        </Link>
        <NavigationMenu>
          <NavigationMenuList>
          <NavigationMenuItem>
              <Link href="/" legacyBehavior passHref>
                <NavigationMenuLink className={cn("px-4 py-2 hover:bg-accent rounded-md")}>
                  Home
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
            <NavigationMenuItem>
              <Link href="/products" legacyBehavior passHref>
                <NavigationMenuLink className={cn("px-4 py-2 hover:bg-accent rounded-md")}>
                  Products
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
            <NavigationMenuItem>
              <Link href="/cart" legacyBehavior passHref>
                <NavigationMenuLink className={cn("px-4 py-2 hover:bg-accent rounded-md")}>
                  Cart
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>
    
        <div className="flex-1 flex justify-end items-center gap-2 mr-1">
          <CartSheet/>
          {!isLoading && (
            <>
              {!isAuthenticated ? (
                <div className="flex gap-2">
                  <Button variant="ghost" asChild>
                    <Link href="/login">Login</Link>
                  </Button>
                  <Button asChild>
                    <Link href="/register">Register</Link>
                  </Button>
                </div>
              ) : (
                <UserMenu />
              )}
            </>
          )}
        </div>
      </div>
      
    </header>
  );
}
