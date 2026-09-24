"use client"

import { Suspense, useEffect, useRef } from "react"
import { useSearchParams } from "next/navigation"
import Link from "next/link"
import { useAuthMutations } from "@/hooks/mutations/use-auth-mutations"

function VerifyEmailContent() {
  const searchParams = useSearchParams()
  const token = searchParams.get("token") ?? ""
  const { verifyEmailMutation } = useAuthMutations()
  const { mutate: verifyEmail } = verifyEmailMutation
  const hasVerified = useRef(false)

  useEffect(() => {
    if (token && !hasVerified.current) {
      hasVerified.current = true
      verifyEmail({ token })
    }
  }, [token, verifyEmail])

  return (
    <div className="min-h-screen flex flex-col justify-center items-center">
      <div className="w-full max-w-2xl bg-white rounded-lg shadow-md px-12 py-16 mx-auto">
        <div className="flex flex-col items-center justify-center gap-1 mb-4">
          {!token && (
            <>
              <span className="text-zinc-900 font-poppins text-lg">
                Invalid link
              </span>
              <span className="text-zinc-500 font-poppins text-base text-center">
                This verification link is invalid.
              </span>
            </>
          )}

          {token && verifyEmailMutation.isPending && (
            <>
              <span className="text-zinc-900 font-poppins text-lg">
                Verifying your email...
              </span>
              <span className="text-zinc-500 font-poppins text-base text-center">
                Please wait a moment.
              </span>
            </>
          )}

          {verifyEmailMutation.isSuccess && (
            <>
              <span className="text-zinc-900 font-poppins text-lg">
                Email verified!
              </span>
              <span className="text-zinc-500 font-poppins text-base text-center">
                Your email has been verified. You can now sign in.
              </span>
            </>
          )}

          {verifyEmailMutation.isError && (
            <>
              <span className="text-zinc-900 font-poppins text-lg">
                Verification failed
              </span>
              <span className="text-zinc-500 font-poppins text-base text-center">
                The link may be expired or invalid. Please request a new verification email.
              </span>
            </>
          )}
        </div>

        <div className="flex justify-center mt-6">
          <Link
            href="/signin"
            className="bg-transparent font-poppins text-xs text-secondary hover:text-hover h-8 px-4 hover:bg-tambo-primary-hover/60 rounded-sm cursor-pointer"
          >
            Back to sign in
          </Link>
        </div>
      </div>
    </div>
  )
}

export default function VerifyEmailPage() {
  return (
    <Suspense>
      <VerifyEmailContent />
    </Suspense>
  )
}
