"use client";
import { User } from "@/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { Label } from "../ui/label";
const profileSchema = z.object({
  name: z.string().min(3),
  email: z.string().email(),
  currentPassword: z.string().min(6).optional(),
  newPassword: z.string().min(6).optional(),
  confirmPassword: z.string().min(6).optional(),
});
// type FormData = z.infer<typeof profileSchema>
export function ProfileForm({ user }: { user: User }) {
  const {
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
  } = useForm({
    resolver: zodResolver(profileSchema),
    defaultValues: {
      name: user?.name || "",
      email: user?.email || "",
      currentPassword: "",
      newPassword: "",
      confirmPassword: "",
    },
  });
  const onSubmit = (data: any) => {
    console.log(data);
  };
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-6 max-w-xl">
      <div className="space-y-4">
        <Label htmlFor="name">Name</Label>
        <Input {...register("name")} error={errors.name?.message} />
        <Label htmlFor="email">Email</Label>
        <Input {...register("email")} error={errors.email?.message} />
      </div>
      <div className="space-y-4">
        <h3 className="font-medium text-lg">Change Password</h3>
       
        <Input
          type="password"
          {...register("currentPassword")}
          error={errors.currentPassword?.message}
          placeholder="Current Password"
        />
        
        <Input
          type="password"
          {...register("newPassword")}
          error={errors.newPassword?.message}
          placeholder="New Password"
        />
        
        <Input
          type="password"
          {...register("confirmPassword")}
          error={errors.confirmPassword?.message}
          placeholder="Confirm Password"
        />
      </div>
      <Button className="w-full mt-4" type="submit">Update Profile</Button>
    </form>
  );
}
