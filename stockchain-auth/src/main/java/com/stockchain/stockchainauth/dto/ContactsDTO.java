package com.stockchain.stockchainauth.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.UUID;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ContactsDTO {
    private UUID id;
    private String key;
    private String value;
}
