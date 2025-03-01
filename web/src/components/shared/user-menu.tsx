"use client";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "../ui/button";
import Link from "next/link";
import { LogOut, Settings, ShoppingBag, User, UserPen, UserRound } from "lucide-react";
export function UserMenu() {
//   const user = {
//     email: "yqZuq@example.com",
//   };
const user= true
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" size="icon">
        <UserRound className="h-5 w-5"/>
        </Button>

      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-56">
        {user ? (
          <>
            <DropdownMenuLabel>{"iangechuki@gmail.com"}</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem asChild>
                {/* UserRound */}
              <Link href="/profile">
              <UserPen className="mr-2 h-4 w-4" />
                Profile
              </Link>
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem asChild>
              <Link href="/orders" className="cursor-pointer">
                <ShoppingBag className="mr-2 h-4 w-4" />
                Orders
              </Link>
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem asChild>
              <Link href="/settings/security" className="cursor-pointer">
                <Settings className="mr-2 h-4 w-4" />
                Settings
              </Link>
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem className="cursor-pointer text-destructive">
              <LogOut className="mr-2 h-4 w-4" />
              Logout
            </DropdownMenuItem>
          </>
        ) : (
          <>
          <DropdownMenuItem asChild>
            <Link href="/login" className="cursor-pointer">
            Login
            </Link>
          </DropdownMenuItem>
          <DropdownMenuItem asChild>
            <Link href="/register" className="cursor-pointer">
            Register
            </Link>
          </DropdownMenuItem>
          </>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
