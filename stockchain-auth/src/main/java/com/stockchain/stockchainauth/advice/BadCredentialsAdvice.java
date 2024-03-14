package com.stockchain.stockchainauth.advice;

import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ProblemDetail;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@ControllerAdvice
public class BadCredentialsAdvice extends ResponseEntityExceptionHandler {
    @ExceptionHandler(value = BadCredentialsException.class)
    protected ResponseEntity<Object> handleUnauthorized(RuntimeException ex, WebRequest request) {
        ProblemDetail problemDetail = ProblemDetail.forStatusAndDetail(
                HttpStatus.UNAUTHORIZED,
                ex.getMessage()
        );
        return handleExceptionInternal(ex, problemDetail, new HttpHeaders(), HttpStatus.UNAUTHORIZED, request);
    }
}