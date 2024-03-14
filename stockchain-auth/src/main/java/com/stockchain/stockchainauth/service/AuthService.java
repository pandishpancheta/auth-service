package com.stockchain.stockchainauth.service;

import com.stockchain.stockchainauth.dto.auth.AuthResponse;
import com.stockchain.stockchainauth.dto.auth.LoginRequest;
import com.stockchain.stockchainauth.dto.auth.RegisterRequest;
import com.stockchain.stockchainauth.dto.auth.RefreshTokenRequest;
import com.stockchain.stockchainauth.entity.User;

public interface AuthService {
    AuthResponse login(LoginRequest loginRequest);
    AuthResponse register(RegisterRequest registerRequest);
    AuthResponse refresh(RefreshTokenRequest refreshTokenRequest);
    User getCurrentUser();
}
