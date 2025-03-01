"use client";
import { Label } from "@/components/ui/label";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import Link from "next/link";
import { useForm } from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {z} from "zod"
import { SocialLogin } from "./social-login";

const loginSchema = z.object({
    email:z.string().email(),
    password:z.string().min(6)
})
type FormData = z.infer<typeof loginSchema>
export function LoginForm() {
const {register,handleSubmit,formState:{errors,isSubmitting}}=useForm<FormData>({
    resolver:zodResolver(loginSchema)
})
const onSubmit = async(data :FormData)=>{
    console.log(data)
}
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="email">Email</Label>
        <Input id="email" type="email" {...register("email")} error={errors.email?.message}/>
      </div>
      <div className="space-y-2">
        <div className="flex items-center">
        <Label htmlFor="password">Password</Label>
        <Link href="/forgot-password" className="ml-auto text-sm underline-offset-4 hover:underline">Forgot password?</Link>
        </div>
        <Input id="password" type="password" {...register("password")} error={errors.password?.message}/>
      </div>
      <Button className="w-full">Sign In</Button>
      <SocialLogin/>
      <div className="text-center text-sm">
   
      </div>
    </form>
  );
}
