export const checkError: (error: any) => Promise<void> = (error) => {
    const status = error.status;
    if (status === 401 || status === 403) {
        return Promise.reject({ redirectTo: '/unauthorized', logoutUser: false });
    }
    // other error code (404, 500, etc): no need to log out
    return Promise.resolve();

}