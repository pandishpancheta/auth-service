package com.stockchain.stockchainauth.advice;

import jakarta.persistence.EntityExistsException;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ProblemDetail;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@ControllerAdvice
public class EntityExistsAdvice extends ResponseEntityExceptionHandler {
    @ExceptionHandler(value = EntityExistsException.class)
    protected ResponseEntity<Object> handleConflict(RuntimeException ex, WebRequest request) {
        ProblemDetail problemDetail = ProblemDetail.forStatusAndDetail(
                HttpStatus.CONFLICT,
                ex.getMessage()
        );
        return handleExceptionInternal(ex, problemDetail, new HttpHeaders(), HttpStatus.CONFLICT, request);
    }
}
