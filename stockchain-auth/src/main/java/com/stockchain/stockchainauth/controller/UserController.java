package com.stockchain.stockchainauth.controller;

import com.stockchain.stockchainauth.dto.UserDTO;
import com.stockchain.stockchainauth.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api/v1/users")
@RequiredArgsConstructor
public class UserController {
    private final UserService userService;

    @GetMapping("/{id}")
    public ResponseEntity<UserDTO> getUserById(UUID id) {
        return ResponseEntity.ok(userService.getUserById(id));
    }

    @GetMapping("/me")
    public ResponseEntity<UserDTO> getCurrentUser() {
        return ResponseEntity.ok(userService.getCurrentUser());
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteUser(UUID id) {
        userService.deleteUser(id);
        return ResponseEntity.ok().build();
    }
}
