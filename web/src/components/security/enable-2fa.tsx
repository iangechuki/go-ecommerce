"use client";

import { useState } from "react";
import { Button } from "../ui/button";
import { QrCode } from "lucide-react";
import { Input } from "../ui/input";
import { InputOTP, InputOTPGroup, InputOTPSlot } from "../ui/input-otp";
import { Dialog } from "@radix-ui/react-dialog";
import {
    DialogClose,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "../ui/dialog";

export function Enable2FA() {
  const [otpauthUrl, setOtpauthUrl] = useState<string | null>(null);
  const [secret, setSecret] = useState("");
  const [message, setMessage] = useState("");
  const [code, setCode] = useState("");
  const [is2FAEnabled, setIs2FAEnabled] = useState(false);
  const [modalOpen, setModalOpen] = useState(false);
  async function handleSetup2FA() {
    const fakeResponse = {
      otpauth_url:
        "otpauth://totp/YourAppName:john.doe@example.com?secret=ABCDEF1234567890&issuer=YourAppName",
      secret: "ABCDEF1234567890",
    };
    setOtpauthUrl(fakeResponse.otpauth_url);
    setSecret(fakeResponse.secret);
    setMessage(
      "Scan the QR code using your authentication app or enter the setup key"
    );
    setModalOpen(true)
  }
  const handleVerify2FA = async () => {
    if (code === "123456") {
        setIs2FAEnabled(true)
        setModalOpen(false)
        setCode("")
      setMessage("2FA is enabled");
      
    }
  };
  const handleDisable2FA = () => {
    setIs2FAEnabled(false);
    setOtpauthUrl(null);
    setSecret("");
    setMessage("");
  };
  return (
    <div className="space-y-4">
      {is2FAEnabled ? (
        <Button variant="destructive" onClick={handleDisable2FA}>
          Disable Two-Factor Authentication
        </Button>
      ) : (
        <Button onClick={handleSetup2FA}>
          Enable Two-Factor Authentication
        </Button>
      )}
      <Dialog open={modalOpen} onOpenChange={setModalOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Enable Two-Factor Authentication</DialogTitle>
            <DialogDescription>{message}</DialogDescription>
          </DialogHeader>
          <div className="flex items-center space-y-4">
            {otpauthUrl && <QrCode size={150} />}
            {secret && <p>{secret}</p>}

            
          </div>
          <input
              type="text"
              placeholder="Enter 2FA Code"
              value={code}
              onChange={(e) => setCode(e.target.value)}
              className="border rounded px-2 py-1 w-full"
            />
            
          <Button onClick={handleVerify2FA} className="w-full">Verify Code</Button>
        </DialogContent>
      </Dialog>
     
    </div>
  );
}
