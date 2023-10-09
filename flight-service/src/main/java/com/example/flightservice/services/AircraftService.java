package com.example.flightservice.services;

import com.example.flightservice.models.Aircraft;
import com.example.flightservice.repositories.AircraftRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Optional;

@Service
public class AircraftService {

    @Autowired
    private AircraftRepository aircraftRepository;

    @Transactional
    public Iterable<Aircraft> getAllAircrafts() {
        return aircraftRepository.findAll();
    }

    @Transactional
    public Aircraft createAircraft(Aircraft aircraft) {
        return aircraftRepository.save(aircraft);
    }

    @Transactional
    public Aircraft getAircraftById(Integer id) {
        return aircraftRepository.findById(id).orElse(null);
    }

    @Transactional
    public Aircraft updateAircraftById(Integer id, Aircraft aircraft) {
        return aircraftRepository.findById(id)
                .map(existingAircraft -> {
                    Optional.ofNullable(aircraft.getModel()).ifPresent(existingAircraft::setModel);
                    Optional.ofNullable(aircraft.getSeatingCapacity()).ifPresent(existingAircraft::setSeatingCapacity);
                    Optional.ofNullable(aircraft.getSeatingConfiguration()).ifPresent(existingAircraft::setSeatingConfiguration);

                    return aircraftRepository.save(existingAircraft);
                })
                .orElse(null);
    }

    @Transactional
    public void deleteAircraftById(Integer id) {
        aircraftRepository.deleteById(id);
    }
}
