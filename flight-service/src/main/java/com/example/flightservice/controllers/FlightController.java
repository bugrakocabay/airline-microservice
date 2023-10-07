package com.example.flightservice.controllers;

import com.example.flightservice.models.Flight;
import com.example.flightservice.services.FlightService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.format.annotation.DateTimeFormat;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.sql.Timestamp;
import java.text.ParseException;
import java.time.LocalDateTime;
import java.util.Date;

@RestController
public class FlightController {

    @Autowired
    private FlightService flightService;

    @GetMapping("/flight")
    public ResponseEntity<Iterable<Flight>> getAllFlights() {
        return new ResponseEntity<>(flightService.getAllFlights(), null, HttpStatus.OK);
    }

    @PostMapping("/flight")
    public ResponseEntity<Flight> createFlight(@RequestBody Flight flight) {
        return new ResponseEntity<>(flightService.createFlight(flight), null, HttpStatus.CREATED);
    }

    @GetMapping("/flight/search")
    public ResponseEntity<Iterable<Flight>> getFlightByArrivalAndDeparture(
            @RequestParam(required = false) String arrival,
            @RequestParam(required = false) String departure,
            @RequestParam(name = "from", required = false) @DateTimeFormat(iso = DateTimeFormat.ISO.DATE_TIME) Date from,
            @RequestParam(name = "to", required = false) @DateTimeFormat(iso = DateTimeFormat.ISO.DATE_TIME) Date to) {
        return new ResponseEntity<>(
                flightService.searchFlight(arrival, departure, from, to),
                null,
                HttpStatus.OK
        );
    }
}
