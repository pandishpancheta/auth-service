package com.stockchain.stockchainauth.service.impl;

import com.stockchain.stockchainauth.dto.UserDTO;
import com.stockchain.stockchainauth.entity.User;
import com.stockchain.stockchainauth.mapper.UserMapper;
import com.stockchain.stockchainauth.repository.UserRepository;
import com.stockchain.stockchainauth.service.AuthService;
import com.stockchain.stockchainauth.service.UserService;
import jakarta.persistence.EntityNotFoundException;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {
    private final UserRepository userRepository;
    private final AuthService authService;

    @Override
    public UserDTO getUserById(UUID id) {
        User user = userRepository.findById(id).orElseThrow(() -> new EntityNotFoundException("Requested user is not available!"));
        return UserMapper.INSTANCE.toUserDTO(user);
    }

    @Override
    public UserDTO getCurrentUser() {
        return UserMapper.INSTANCE.toUserDTO(authService.getCurrentUser());
    }

    @Override
    public void deleteUser() {
        User user = authService.getCurrentUser();
        userRepository.delete(user);
    }
}
