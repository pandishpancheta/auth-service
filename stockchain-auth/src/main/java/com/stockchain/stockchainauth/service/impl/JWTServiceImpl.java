package com.stockchain.stockchainauth.service.impl;

import com.stockchain.stockchainauth.entity.User;
import com.stockchain.stockchainauth.service.JWTService;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.io.Decoders;
import io.jsonwebtoken.security.Keys;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Service;

import java.security.Key;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.function.Function;

@Service
public class JWTServiceImpl implements JWTService {
    private final String SECRET_KEY;
    private final long ACCESS_TOKEN_EXPIRATION = 900000;
    private final long REFRESH_TOKEN_EXPIRATION = 604800000;


    public JWTServiceImpl(@Value("${stockchain.app.jwt-secret}") String secretKey) {
        this.SECRET_KEY = secretKey;
    }

    @Override
    public String extractUsername(String jwt) {
        return extractClaim(jwt, Claims::getSubject);
    }

    @Override
    public <T> T extractClaim(String jwt, Function<Claims, T> claimsResolver) {
        final Claims claims = extractAllClaims(jwt);
        return claimsResolver.apply(claims);
    }

    @Override
    public String generateToken(UserDetails userDetails, Boolean access) {
        Date expiration;
        if (access) {
            expiration = new Date(System.currentTimeMillis() + ACCESS_TOKEN_EXPIRATION);
        } else {
            expiration = new Date(System.currentTimeMillis() + REFRESH_TOKEN_EXPIRATION);
        }
        return generateToken(new HashMap<>(), userDetails, expiration);
    }

    @Override
    public String generateToken(Map<String, Object> extractClaims, UserDetails userDetails, Date expiration) {
        return Jwts.builder()
                .setClaims(extractClaims)
                .setSubject(((User) userDetails).getEmail())
                .setIssuedAt(new Date(System.currentTimeMillis()))
                .setExpiration(expiration)
                .signWith(getSignInKey(), SignatureAlgorithm.HS256)
                .compact();
    }


    @Override
    public boolean isTokenValid(String jwt, UserDetails userDetails) {
        final String username = extractUsername(jwt);
        return (username.equals(((User) userDetails).getEmail())) && !isTokenExpired(jwt);
    }

    private boolean isTokenExpired(String jwt) {
        return extractExpiration(jwt).before(new Date());
    }

    private Date extractExpiration(String jwt) {
        return extractClaim(jwt, Claims::getExpiration);
    }

    private Claims extractAllClaims(String jwt) {
        return Jwts.parserBuilder()
                .setSigningKey(getSignInKey())
                .build()
                .parseClaimsJws(jwt)
                .getBody();
    }

    private Key getSignInKey() {
        byte[] keyBytes = Decoders.BASE64.decode(SECRET_KEY);
        return Keys.hmacShaKeyFor(keyBytes);
    }
}