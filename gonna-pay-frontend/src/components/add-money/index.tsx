"use client";

import { useEffect, useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { InitialHelper } from "@/components/initial-helper";
import { BanknoteArrowUp, Loader2 } from "lucide-react";
import { Button } from "../ui/button";
import QRCode from "react-qr-code";
import { toast } from "sonner";

export const AddMoney = () => {
  const [amount, setAmount] = useState("");

  const [modal, setModal] = useState({
    isOpen: false,
    step: 1,
    isLoading: false,
  });

  const [timer, setTimer] = useState(60);

  const isTimerActive = modal.step === 2 && modal.isOpen;
  const isTimerVisible = modal.step === 2;

  useEffect(() => {
    if (isTimerActive) {
      setTimer(60);
      const interval = setInterval(() => {
        setTimer((prev) => {
          if (prev <= 1) {
            clearInterval(interval);
            toast.error("Time expired! Generate a new PIX.");
            setModal((prev) => ({ ...prev, step: 1 }));
            return 0;
          }
          return prev - 1;
        });
      }, 1000);

      return () => clearInterval(interval);
    }
  }, [isTimerActive]);

  const pixCode = `00020126580014br.gov.bcb.pix0136${amount}52040000530398654${amount.padStart(
    10,
    "0"
  )}5802BR5925Your Name Here6009SAO PAULO`;

  const handleContinue = async () => {
    if (!amount || parseFloat(amount) <= 0) return;

    setModal((prev) => ({ ...prev, isLoading: true }));
    await new Promise((resolve) => setTimeout(resolve, 1500));
    setModal((prev) => ({ ...prev, step: 2, isLoading: false }));
  };

  const handleFinish = async () => {
    setModal({ isOpen: false, step: 1, isLoading: true });

    toast.success(`Successfully added $${amount} to your wallet!`);

    await new Promise((resolve) => setTimeout(resolve, 1500));

    setModal({
      isOpen: false,
      step: 1,
      isLoading: false,
    });

    setAmount("");
  };

  const handleCancel = () => {
    setModal({ isOpen: false, step: 1, isLoading: false });

    setAmount("");
  };

  return (
    <Dialog
      open={modal.isOpen}
      onOpenChange={(isOpen) => setModal((prev) => ({ ...prev, isOpen }))}
    >
      <DialogTrigger>
        <InitialHelper name="Add money" icon={BanknoteArrowUp} />
      </DialogTrigger>

      <DialogContent>
        {modal.step === 1 && (
          <>
            <DialogHeader>
              <DialogTitle>Add Money</DialogTitle>
              <DialogDescription>
                How much would you like to add to your wallet?
              </DialogDescription>
            </DialogHeader>

            <div className="flex flex-col gap-4">
              <input
                type="number"
                placeholder="$ 0,00"
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                className="border rounded-md px-3 py-2 text-lg"
                min="0"
                step="0.01"
              />

              <div className="flex flex-row gap-2">
                <Button
                  className="w-1/2"
                  variant="ghost"
                  onClick={handleCancel}
                >
                  Cancel
                </Button>
                <Button
                  className="w-1/2 transition-all duration-150 ease-in-out"
                  variant="outline"
                  onClick={handleContinue}
                  disabled={
                    modal.isLoading || !amount || parseFloat(amount) <= 0
                  }
                >
                  {modal.isLoading ? (
                    <>
                      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    </>
                  ) : (
                    "Continue"
                  )}
                </Button>
              </div>
            </div>
          </>
        )}

        {modal.step === 2 && (
          <>
            <DialogHeader>
              <DialogTitle>Scan your QR Code</DialogTitle>
              <DialogDescription>
                Use seu app de pagamento para escanear o código abaixo e
                finalizar a transação de $ {amount}.
              </DialogDescription>
            </DialogHeader>

            <div className="flex flex-col items-center gap-4 py-4">
              <div className="bg-white p-4 rounded-lg">
                <QRCode value={pixCode} size={200} />
              </div>

              {isTimerVisible &&
                (() => {
                  const isExpiring = timer <= 10;
                  const colorClass = isExpiring
                    ? "bg-red-50 border-red-200 text-red-600"
                    : "bg-green-50 border-green-200 text-green-600";

                  return (
                    <div
                      className={`flex items-center gap-2 px-4 py-2 rounded-lg border transition-all duration-300 ${colorClass}`}
                    >
                      <p className="text-sm font-medium">
                        Time remaining: {timer}s
                      </p>
                    </div>
                  );
                })()}

              <div className="w-full">
                <p className="text-xs text-zinc-500 mb-1">
                  Or copy the QR code:
                </p>
                <div className="flex gap-2">
                  <input
                    type="text"
                    value={pixCode}
                    readOnly
                    className="flex-1 border rounded-md px-3 py-2 text-xs bg-zinc-50"
                  />
                  <Button
                    size="sm"
                    variant="outline"
                    onClick={() => {
                      navigator.clipboard.writeText(pixCode);
                      toast.success("copied successfully!");
                    }}
                  >
                    Copy
                  </Button>
                </div>
              </div>

              <Button
                className="w-full"
                variant="outline"
                onClick={handleFinish}
                disabled={modal.isLoading}
              >
                {modal.isLoading ? (
                  <Loader2 className="h-4 w-4 animate-spin" />
                ) : (
                  "Finish"
                )}
              </Button>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};
