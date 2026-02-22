package com.securerent.riskassessment.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import java.time.LocalDateTime;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class RiskAssessmentResult {
    private String renterId;
    private int riskScore;
    private boolean qualified;
    private double maxCoverage;
    private String recommendation;
    private LocalDateTime evaluatedAt;
}
