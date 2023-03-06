import 'package:equatable/equatable.dart';
import 'package:flutter/material.dart';
import 'package:journey/core/models/state.dart';

@immutable
abstract class EntityStateState extends Equatable {
  const EntityStateState();
}

class EntityStateLoading extends EntityStateState {
  @override
  List<Object> get props => [];
}

class EntityStateUpdated extends EntityStateState {
  final EntityState state;

  const EntityStateUpdated({ required this.state });

  @override
  List<Object> get props => [state];
}

class EntityStateError extends EntityStateState {
  @override
  List<Object> get props => [];
}