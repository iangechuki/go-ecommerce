"use client";
import { OrderList, OrderListSkeleton } from "@/components/orders/order-list";
import { ProfileForm } from "@/components/profile/profile-form";
import { ThemeToggle } from "@/components/theme-toggle";
import { Skeleton } from "@/components/ui/skeleton";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { orders } from "@/constants/orders";

export default function ProfilePage() {
  const user = {
    id: "1",
    email: "iangechuki@gmail.com",
    name: "ian ochako",
    password:"password",
    role:"user"
  };
  const isLoading = false;
  return (
    <div className="container py-8">
      <div className="max-w-4xl mx-auto space-y-4">
        <div className="flex justify-between items-start">
          <div>
            <h1 className="text-3xl font-bold">
              {isLoading ? (
                <Skeleton className="h-8 w-48" />
              ) : (
                `Welcome ${user?.name}`
              )}
            </h1>
            <p className="text-muted-foreground">
              {isLoading ? <Skeleton className="h-4 w-64 mt-2" /> : user?.email}
            </p>
          </div>
          <ThemeToggle/>
        </div>
        <Tabs defaultValue="orders">
          <TabsList>
            <TabsTrigger value="orders">Orders</TabsTrigger>
            <TabsTrigger value="settings">Settings</TabsTrigger>
            </TabsList>
            <TabsContent value="orders" asChild>
            {isLoading ? (
              <OrderListSkeleton />
            ) : (
              <OrderList orders={orders} />
            )}
            </TabsContent>
            <TabsContent value="settings">
            <ProfileForm user={user}/>
          </TabsContent>
        </Tabs>
      </div>
    </div>
  );
}
