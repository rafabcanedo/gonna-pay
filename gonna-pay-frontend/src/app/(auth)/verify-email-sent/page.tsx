import Link from "next/link"

export default function VerifyEmailSentPage() {
  return (
    <div className="min-h-screen flex flex-col justify-center items-center">
      <div className="w-full max-w-2xl bg-white rounded-lg shadow-md px-12 py-16 mx-auto">
        <div className="flex flex-col items-center justify-center gap-1 mb-4">
          <span className="text-zinc-900 font-poppins text-lg">
            Check your inbox
          </span>
          <span className="text-zinc-500 font-poppins text-base text-center">
            We sent a verification link to your email. Click the link to verify your account.
          </span>
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
