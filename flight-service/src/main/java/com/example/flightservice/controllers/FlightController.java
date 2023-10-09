package com.example.flightservice.controllers;

import com.example.flightservice.dtos.responses.BaseResponse;
import com.example.flightservice.models.Flight;
import com.example.flightservice.services.FlightService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.format.annotation.DateTimeFormat;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Date;

@RestController
public class FlightController {

    @Autowired
    private FlightService flightService;

    @GetMapping("/flight")
    public ResponseEntity<BaseResponse<Iterable<Flight>>> getAllFlights() {
        BaseResponse<Iterable<Flight>> response = BaseResponse.<Iterable<Flight>>builder()
                .error(false)
                .message("OK")
                .data(flightService.getAllFlights())
                .build();

        return new ResponseEntity<>(response, null, HttpStatus.OK);
    }

    @PostMapping("/flight")
    public ResponseEntity<BaseResponse<Flight>> createFlight(@RequestBody Flight flight) {
        BaseResponse<Flight> response = BaseResponse.<Flight>builder()
                .error(false)
                .message("OK")
                .data(flightService.createFlight(flight))
                .build();

        return new ResponseEntity<>(response, null, HttpStatus.CREATED);
    }

    @GetMapping("/flight/search")
    public ResponseEntity<BaseResponse<Iterable<Flight>>> getFlightByArrivalAndDeparture(
            @RequestParam(required = false) String arrival,
            @RequestParam(required = false) String departure,
            @RequestParam(name = "from", required = false) @DateTimeFormat(iso = DateTimeFormat.ISO.DATE_TIME) Date from,
            @RequestParam(name = "to", required = false) @DateTimeFormat(iso = DateTimeFormat.ISO.DATE_TIME) Date to) {
        BaseResponse<Iterable<Flight>> response = BaseResponse.<Iterable<Flight>>builder()
                .error(false)
                .message("OK")
                .data(flightService.searchFlight(arrival, departure, from, to))
                .build();

        return new ResponseEntity<>(response, null, HttpStatus.OK);
    }

    @GetMapping("/flight/{id}")
    public ResponseEntity<BaseResponse<Flight>> getFlightById(@PathVariable Integer id) {
        Flight flight = flightService.getFlightById(id);

        if (flight == null) {
            BaseResponse<Flight> response = BaseResponse.<Flight>builder()
                    .error(true)
                    .message("Flight not found")
                    .build();
            return new ResponseEntity<>(response, null, HttpStatus.NOT_FOUND);
        }

        BaseResponse<Flight> response = BaseResponse.<Flight>builder()
                .error(false)
                .message("OK")
                .data(flight)
                .build();
        return new ResponseEntity<>(response, null, HttpStatus.OK);
    }

    @PutMapping("/flight/{id}")
    public ResponseEntity<BaseResponse<Flight>> updateFlightById(@PathVariable Integer id, @RequestBody Flight flight) {
        Flight updatedFlight = flightService.updateFlightById(id, flight);

        if (updatedFlight == null) {
            BaseResponse<Flight> response = BaseResponse.<Flight>builder()
                    .error(true)
                    .message("Flight not found")
                    .build();
            return new ResponseEntity<>(response, null, HttpStatus.NOT_FOUND);
        }

        BaseResponse<Flight> response = BaseResponse.<Flight>builder()
                .error(false)
                .message("OK")
                .data(updatedFlight)
                .build();
        return new ResponseEntity<>(response, null, HttpStatus.OK);
    }

    @DeleteMapping("/flight/{id}")
    public ResponseEntity<BaseResponse<Flight>> deleteFlightById(@PathVariable Integer id) {
        Flight flight = flightService.getFlightById(id);

        if (flight == null) {
            BaseResponse<Flight> response = BaseResponse.<Flight>builder()
                    .error(true)
                    .message("Flight not found")
                    .build();
            return new ResponseEntity<>(response, null, HttpStatus.NOT_FOUND);
        }

        flightService.deleteFlightById(id);

        BaseResponse<Flight> response = BaseResponse.<Flight>builder()
                .error(false)
                .message("OK")
                .data(flight)
                .build();
        return new ResponseEntity<>(response, null, HttpStatus.OK);
    }
}
