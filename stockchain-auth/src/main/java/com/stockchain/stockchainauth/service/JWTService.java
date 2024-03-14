package com.stockchain.stockchainauth.service;

import io.jsonwebtoken.Claims;
import org.springframework.security.core.userdetails.UserDetails;

import java.util.Date;
import java.util.Map;
import java.util.function.Function;

public interface JWTService {
    String extractUsername(String jwt);
    <T> T extractClaim(String jwt, Function<Claims, T> claimsResolver);
    String generateToken(UserDetails userDetails, Boolean access);
    String generateToken(Map<String, Object> extractClaims, UserDetails userDetails, Date expiration);
    boolean isTokenValid(String jwt, UserDetails userDetails);
}