'use client'

export const GeneralError = ({ error, reset }: { error: Error & { digest?: string }; reset: () => void }) => {
    return (
        <div className="flex items-center justify-center min-h-screen">
            <div className="p-6 bg-red-50 border border-red-200 rounded-lg max-w-md">
                <h2 className="text-lg font-semibold text-red-600 mb-2">
                    Something went wrong
                </h2>
                <p className="text-red-600 text-sm mb-4">
                    {error.message || 'An unexpected error occurred'}
                </p>
                <button
                    onClick={() => reset()}
                    className="w-full px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition"
                >
                    Try again
                </button>
            </div>
        </div>
    )
}