package com.securerent.riskassessment.controller;

import com.securerent.riskassessment.model.RenterProfile;
import com.securerent.riskassessment.model.RiskAssessmentResult;
import com.securerent.riskassessment.service.RiskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@RestController
@RequestMapping("/api/v1")
public class RiskController {

    @Autowired
    private RiskService riskService;

    @PostMapping("/assess")
    public ResponseEntity<RiskAssessmentResult> assess(@RequestBody RenterProfile profile) {
        return ResponseEntity.ok(riskService.assessRenter(profile));
    }

    @GetMapping("/health")
    public ResponseEntity<Map<String, String>> healthCheck() {
        return ResponseEntity.ok(Map.of("status", "up"));
    }
}
