import { useForm } from "react-hook-form";
import { Button } from "../ui/button";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
const forgotPasswordSchema = z.object({
  email: z.string().email(),
});
type FormData = z.infer<typeof forgotPasswordSchema>;

export function ForgotPasswordForm() {
  const {
    handleSubmit,
    register,
    reset,
    formState: { isSubmitting, errors },
  } = useForm({
    resolver: zodResolver(forgotPasswordSchema),
  });
  const onSubmit = (data: FormData) => {
    console.log(data)
    reset()
  };
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="email">Email</Label>
        <Input
          id="email"
          type="email"
          {...register("email")}
          error={errors.email?.message}
        />
      </div>
      <Button type="submit" className="w-full" loading={isSubmitting}>
        Send Reset Link
      </Button>
    </form>
  );
}
