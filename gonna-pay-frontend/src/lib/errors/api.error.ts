export class ApiError extends Error {
  constructor(
    public status: number,
    public message: string,
    public details?: unknown
  ) {
    super(message)
    this.name = 'ApiError'
  }
}

export class ValidationError extends ApiError {
  constructor(public fields: Record<string, string[]>) {
    super(400, 'Validation Failed', fields)
    this.name = 'ValidationError'
  }
}

export class UnauthorizedError extends ApiError {
  constructor() {
    super(401, 'Unauthorized. Please login.', null)
    this.name = 'UnauthorizedError'
  }
}

export class ForbiddenError extends ApiError {
  constructor() {
    super(403, 'You do not have permission to access this resource.', null)
    this.name = 'ForbiddenError'
  }
}

export class NotFoundError extends ApiError {
  constructor(resource: string) {
    super(404, `${resource} not found`, null)
    this.name = 'NotFoundError'
  }
}

export class InternalServerError extends ApiError {
  constructor() {
    super(500, 'Internal server error. Please try again later.', null)
    this.name = 'InternalServerError'
  }
}

export class ServiceUnavailableError extends ApiError {
  constructor() {
    super(503, 'Service temporarily unavailable. Please try again later.', null)
    this.name = 'ServiceUnavailableError'
  }
}

export class NetworkError extends ApiError {
  constructor(message = 'Network error. Check your connection.') {
    super(0, message, null)
    this.name = 'NetworkError'
  }
}