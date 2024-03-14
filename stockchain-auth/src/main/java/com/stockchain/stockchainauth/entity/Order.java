package com.stockchain.stockchainauth.entity;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.UUID;

@Entity
@Data
@AllArgsConstructor
@NoArgsConstructor
@Table(
        name = "`order`",
        indexes = {
                @Index(name = "order_listingId_index", columnList = "id", unique = true),
                @Index(name = "order_listingId_index", columnList = "listingId", unique = false)
        }
)
public class Order {
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    @Column(nullable = false)
    private UUID listingId;

    @Column(nullable = false)
    private String token_uri;

    @Column(nullable = false)
    private String status;
}
