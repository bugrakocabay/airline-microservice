package com.example.flightservice.repositories;

import com.example.flightservice.models.Aircraft;
import org.springframework.data.repository.CrudRepository;

public interface AircraftRepository extends CrudRepository<Aircraft, Integer> {
}
