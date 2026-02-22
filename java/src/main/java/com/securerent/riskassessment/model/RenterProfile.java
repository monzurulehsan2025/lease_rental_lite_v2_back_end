package com.securerent.riskassessment.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import java.time.LocalDateTime;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class RenterProfile {
    private String id;
    private String firstName;
    private String lastName;
    private String email;
    private double annualIncome;
    private int creditScore;
    private double monthlyRent;
    private String employmentStatus;
    private LocalDateTime createdAt;
}
