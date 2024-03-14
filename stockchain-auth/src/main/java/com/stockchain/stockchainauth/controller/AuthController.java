package com.stockchain.stockchainauth.controller;

import com.stockchain.stockchainauth.dto.auth.AuthResponse;
import com.stockchain.stockchainauth.dto.auth.LoginRequest;
import com.stockchain.stockchainauth.dto.auth.RegisterRequest;
import com.stockchain.stockchainauth.dto.auth.RefreshTokenRequest;
import com.stockchain.stockchainauth.service.AuthService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/auth")
@RequiredArgsConstructor
public class AuthController {
    private final AuthService authService;

    @PostMapping("/register")
    public ResponseEntity<AuthResponse> register(@RequestBody RegisterRequest registerRequest) {
        return ResponseEntity.ok(authService.register(registerRequest));
    }

    @PostMapping("/login")
    public ResponseEntity<AuthResponse> login(@RequestBody LoginRequest loginRequest) {
        return ResponseEntity.ok(authService.login(loginRequest));
    }

    @PostMapping("/refreshToken")
    public ResponseEntity<AuthResponse> refresh(@RequestBody RefreshTokenRequest refreshTokenRequest) {
        return ResponseEntity.ok(authService.refresh(refreshTokenRequest));
    }
}
