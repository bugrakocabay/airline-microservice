package com.example.flightservice.controllers;

import com.example.flightservice.dtos.responses.BaseResponse;
import com.example.flightservice.models.Aircraft;
import com.example.flightservice.services.AircraftService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
public class AircraftController {

    @Autowired
    private AircraftService aircraftService;

    @GetMapping("/aircraft")
    public ResponseEntity<BaseResponse<Iterable<Aircraft>>> getAllAircrafts() {
        BaseResponse<Iterable<Aircraft>> response = BaseResponse.<Iterable<Aircraft>>builder()
                .error(false)
                .message("OK")
                .data(aircraftService.getAllAircrafts())
                .build();

        return new ResponseEntity<>(response, null, 200);
    }

    @GetMapping("/aircraft/{id}")
    public ResponseEntity<BaseResponse<Aircraft>> getAircraftById(@PathVariable Integer id) {
        Aircraft aircraft = aircraftService.getAircraftById(id);

        BaseResponse<Aircraft> response;

        if (aircraft == null) {
            response = BaseResponse.<Aircraft>builder()
                    .error(true)
                    .message("Aircraft not found")
                    .build();
            return new ResponseEntity<>(response, HttpStatus.NOT_FOUND);
        }

        response = BaseResponse.<Aircraft>builder()
                .error(false)
                .message("OK")
                .data(aircraft)
                .build();
        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    @PostMapping("/aircraft")
    public ResponseEntity<BaseResponse<Aircraft>> createAircraft(@RequestBody Aircraft aircraft) {
        BaseResponse<Aircraft> response = BaseResponse.<Aircraft>builder()
                .error(false)
                .message("OK")
                .data(aircraftService.createAircraft(aircraft))
                .build();

        return new ResponseEntity<>(response, null, 201);
    }

    @PutMapping("/aircraft/{id}")
    public ResponseEntity<BaseResponse<Aircraft>> updateAircraftById(@PathVariable Integer id, @RequestBody Aircraft aircraft) {
        Aircraft existingAircraft = aircraftService.getAircraftById(id);

        BaseResponse<Aircraft> response;

        if (existingAircraft == null) {
            response = BaseResponse.<Aircraft>builder()
                    .error(true)
                    .message("Aircraft not found")
                    .build();
            return new ResponseEntity<>(response, HttpStatus.NOT_FOUND);
        }

        Aircraft updatedAircraft = aircraftService.updateAircraftById(id, aircraft);

        response = BaseResponse.<Aircraft>builder()
                .error(false)
                .message("OK")
                .data(updatedAircraft)
                .build();
        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    @DeleteMapping("/aircraft/{id}")
    public ResponseEntity<BaseResponse<Aircraft>> deleteAircraftById(@PathVariable Integer id) {
        Aircraft existingAircraft = aircraftService.getAircraftById(id);

        BaseResponse<Aircraft> response;

        if (existingAircraft == null) {
            response = BaseResponse.<Aircraft>builder()
                    .error(true)
                    .message("Aircraft not found")
                    .build();
            return new ResponseEntity<>(response, HttpStatus.NOT_FOUND);
        }

        aircraftService.deleteAircraftById(id);

        response = BaseResponse.<Aircraft>builder()
                .error(false)
                .message("OK")
                .build();
        return new ResponseEntity<>(response, HttpStatus.OK);
    }
}
