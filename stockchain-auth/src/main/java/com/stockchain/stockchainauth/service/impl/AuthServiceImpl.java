package com.stockchain.stockchainauth.service.impl;

import com.stockchain.stockchainauth.dto.auth.AuthResponse;
import com.stockchain.stockchainauth.dto.auth.LoginRequest;
import com.stockchain.stockchainauth.dto.auth.RegisterRequest;
import com.stockchain.stockchainauth.dto.auth.RefreshTokenRequest;
import com.stockchain.stockchainauth.entity.User;
import com.stockchain.stockchainauth.exception.InvalidTokenException;
import com.stockchain.stockchainauth.mapper.AuthMapper;
import com.stockchain.stockchainauth.repository.UserRepository;
import com.stockchain.stockchainauth.service.AuthService;
import com.stockchain.stockchainauth.service.JWTService;
import jakarta.persistence.EntityExistsException;
import lombok.RequiredArgsConstructor;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class AuthServiceImpl implements AuthService {
    private final UserRepository userRepository;
    private final JWTService jwtService;
    private final AuthenticationManager authenticationManager;
    private final PasswordEncoder passwordEncoder;

    @Override
    public AuthResponse login(LoginRequest loginRequest) {
        User user;
        try {
            Authentication auth = authenticationManager.authenticate(
                    new UsernamePasswordAuthenticationToken(
                            loginRequest.getEmail(),
                            loginRequest.getPassword()
                    )
            );
            user = (User) auth.getPrincipal();
        } catch (AuthenticationException e) {
            throw new BadCredentialsException("Email and password do not match!");
        }

        String jwt = jwtService.generateToken(user, true);
        String refresh = jwtService.generateToken(user, false);

        return AuthResponse.builder()
                .token(jwt)
                .refresh(refresh)
                .build();
    }

    @Override
    public AuthResponse register(RegisterRequest registerRequest) {
        User user = AuthMapper.INSTANCE.fromRegisterDTO(registerRequest);
        user.setPassword(passwordEncoder.encode(registerRequest.getPassword()));

        try {
            userRepository.save(user);
        } catch (DataIntegrityViolationException e) {
            throw new EntityExistsException("User with the specified credentials already exists!");
        }

        String jwt = jwtService.generateToken(user, true);
        String refresh = jwtService.generateToken(user, false);

        return AuthResponse.builder()
                .token(jwt)
                .refresh(refresh)
                .build();
    }

    @Override
    public AuthResponse refresh(RefreshTokenRequest refreshTokenRequest) {
        try {
            String email = jwtService.extractUsername(refreshTokenRequest.getToken());
            User user = userRepository.findByEmail(email)
                    .orElseThrow(() -> new InvalidTokenException("Invalid Token!"));

            if (jwtService.isTokenValid(refreshTokenRequest.getToken(), user)) {
                String jwt = jwtService.generateToken(user, true);

                return AuthResponse.builder()
                        .token(jwt)
                        .refresh(refreshTokenRequest.getToken())
                        .build();
            } else {
                throw new InvalidTokenException("Invalid Token!");
            }
        } catch (InvalidTokenException ex) {
            throw ex;
        } catch (Exception ex) {
            throw new RuntimeException("An error occurred while refreshing token!", ex);
        }
    }

}
