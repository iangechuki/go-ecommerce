import { LoginForm } from "@/components/auth/login-form";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import Link from "next/link";

export default function LoginPage(){
    return (
        <div className="min-h-screen flex items-center justify-center">
            <Card className="w-full max-w-md">
                <CardHeader>
                    <CardTitle>Sign In</CardTitle>
                </CardHeader>
                <CardContent>
                    <LoginForm/>
                    <div className="text-center text-sm">
                        Don't have an account?
                        <Link href="/register" className="underline">
                            Sign Up</Link>
                    </div>
                </CardContent>
            </Card>

        </div>
    )
}