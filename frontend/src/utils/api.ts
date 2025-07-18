const baseURL = import.meta.env.VITE_API_BASE_URL || "";

interface ApiOptions extends RequestInit {
  requiresAuth?: boolean;
}

export const apiRequest = async (
  endpoint: string,
  options: ApiOptions = {}
): Promise<Response> => {
  const { requiresAuth = false, headers = {}, ...restOptions } = options;
  
  const requestHeaders: Record<string, string> = {
    "Content-Type": "application/json",
    ...headers as Record<string, string>,
  };

  if (requiresAuth) {
    const token = localStorage.getItem("token");
    if (token) {
      requestHeaders.Authorization = `Bearer ${token}`;
    }
  }

  const response = await fetch(`${baseURL}${endpoint}`, {
    ...restOptions,
    headers: requestHeaders,
    credentials: 'include',
  });

  return response;
};

export const get = async (endpoint: string, requiresAuth: boolean = false): Promise<Response> => {
  return apiRequest(endpoint, { method: "GET", requiresAuth });
};

export const post = async (
  endpoint: string,
  data?: unknown,
  requiresAuth: boolean = false
): Promise<Response> => {
  return apiRequest(endpoint, {
    method: "POST",
    body: data ? JSON.stringify(data) : undefined,
    requiresAuth,
  });
};

export const put = async (
  endpoint: string,
  data?: unknown,
  requiresAuth: boolean = false
): Promise<Response> => {
  return apiRequest(endpoint, {
    method: "PUT",
    body: data ? JSON.stringify(data) : undefined,
    requiresAuth,
  });
};

export const del = async (endpoint: string, requiresAuth: boolean = false): Promise<Response> => {
  return apiRequest(endpoint, { method: "DELETE", requiresAuth });
};