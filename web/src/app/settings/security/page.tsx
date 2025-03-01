import { Enable2FA } from "@/components/security/enable-2fa";
import { SecurityForm } from "@/components/security/security-form";
import { SessionList } from "@/components/security/session-list-form";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export default function SecuritySettingsPage() {
  return (
    <div className="flex flex-col items-center gap-4 p-4">
      <Card className="w-full max-w-2xl">
        <CardHeader>
          <CardTitle>Settings</CardTitle>
        </CardHeader>
        <CardContent>
          <SecurityForm/>
        </CardContent>
      </Card>
      <Card className="w-full max-w-2xl">
        <CardHeader>
          <CardTitle> Active Sessions</CardTitle>
        </CardHeader>
        <CardContent><SessionList/></CardContent>
      </Card>
      <Card className="w-full max-w-2xl">
        <CardHeader>
          <CardTitle>Two-Factor Authentication</CardTitle>
        </CardHeader>
        <CardContent><Enable2FA/></CardContent>
      </Card>
    </div>
  );
}
