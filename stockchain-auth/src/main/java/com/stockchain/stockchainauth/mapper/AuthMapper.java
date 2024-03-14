package com.stockchain.stockchainauth.mapper;

import com.stockchain.stockchainauth.dto.auth.AuthResponse;
import com.stockchain.stockchainauth.dto.auth.RegisterRequest;
import com.stockchain.stockchainauth.entity.User;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.factory.Mappers;

@Mapper
public interface AuthMapper {
    AuthMapper INSTANCE = Mappers.getMapper(AuthMapper.class);

    @Mapping(target = "password", ignore = true)
    User fromRegisterDTO(RegisterRequest registerRequest);
}
