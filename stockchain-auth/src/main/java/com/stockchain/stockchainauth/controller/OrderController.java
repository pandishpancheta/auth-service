package com.stockchain.stockchainauth.controller;

import com.stockchain.stockchainauth.dto.OrderDTO;
import com.stockchain.stockchainauth.service.OrderService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api/v1/orders")
@RequiredArgsConstructor
public class OrderController {
    private final OrderService orderService;

    @GetMapping
    public ResponseEntity<List<OrderDTO>> getOrders() {
        return ResponseEntity.ok(orderService.getAllOrders());
    }

    @GetMapping("/{id}")
    public ResponseEntity<OrderDTO> getOrderById(UUID id) {
        return ResponseEntity.ok(orderService.getOrderById(id));
    }

    @PostMapping
    public ResponseEntity<OrderDTO> createOrder(@RequestBody OrderDTO orderDTO) {
        return ResponseEntity.ok(orderService.createOrder(orderDTO));
    }

    @PutMapping("/{id}")
    public ResponseEntity<OrderDTO> updateOrder(@PathVariable UUID id, @RequestBody OrderDTO orderDTO) {
        return ResponseEntity.ok(orderService.updateOrder(id, orderDTO));
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteOrder(@PathVariable UUID id) {
        orderService.deleteOrder(id);
        return ResponseEntity.ok().build();
    }
}
